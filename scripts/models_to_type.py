# --------------------------------------------------------------------------------------------------
#  COPYRIGHT(c) 2023 MONTA                                                                         -
#                                                                                                  -
#  This file is part of Monta.                                                                     -
#                                                                                                  -
#  The Monta software is licensed under the Business Source License 1.1. You are granted the right -
#  to copy, modify, and redistribute the software, but only for non-production use or with a total -
#  of less than three server instances. Starting from the Change Date (November 16, 2026), the     -
#  software will be made available under version 2 or later of the GNU General Public License.     -
#  If you use the software in violation of this license, your rights under the license will be     -
#  terminated automatically. The software is provided "as is," and the Licensor disclaims all      -
#  warranties and conditions. If you use this license's text or the "Business Source License" name -
#  and trademark, you must comply with the Licensor's covenants, which include specifying the      -
#  Change License as the GPL Version 2.0 or a compatible license, specifying an Additional Use     -
#  Grant, and not modifying the license in any other way.                                          -
# --------------------------------------------------------------------------------------------------
import ast
import argparse
import os
import rich
from rich.progress import Progress

# Define the mapping from Django fields to TypeScript types
TYPE_MAP = {
    "UUIDField": "string",
    "OneToOneField": "string",
    "PositiveIntegerField": "number",
    "DecimalField": "number",
    "BooleanField": "boolean",
    "CharField": "string",
    "TextField": "string",
    "DateField": "Date",
    "ChoiceField": "string",
    "ForeignKey": "string",
    "ManyToManyField": "string[]",
    "JSONField": "any",
    "AutoField": "number",
    "BigIntegerField": "number",
    "BinaryField": "any",
    "DateTimeField": "Date",
    "DurationField": "any",
    "EmailField": "string",
    "FileField": "any",
    "FilePathField": "any",
    "FloatField": "number",
    "GenericIPAddressField": "string",
    "ImageField": "any",
    "IntegerField": "number",
    "IPAddressField": "string",
    "NullBooleanField": "boolean",
    "PositiveSmallIntegerField": "number",
}


class ModelVisitor(ast.NodeVisitor):
    """An AST node visitor that finds Django model classes in a file.

    Attributes:
        models: A dictionary where keys are model class names and values are
            lists of tuples, each containing a field name and its corresponding
            TypeScript type.
    """

    def __init__(self):
        """
        Initialize the node visitor with an empty models dictionary.
        """
        self.models = {}

    def visit_ClassDef(self, node: ast.ClassDef) -> None:
        """Visit a ClassDef node in the AST. If the class contains Django model
        fields, add it to the models dictionary.

        Args:
            node(ast.ClassDef): An AST ClassDef node to visit.

        Returns: this function does not return anything.
        """
        fields = []

        for stmt in node.body:
            if isinstance(stmt, ast.Assign):
                for target in stmt.targets:
                    if isinstance(target, ast.Name):
                        field_name = target.id

                        if isinstance(stmt.value, ast.Call):
                            if isinstance(stmt.value.func, ast.Attribute):
                                field_type = stmt.value.func.attr
                            elif isinstance(stmt.value.func, ast.Name):
                                field_type = stmt.value.func.id

                            if ts_type := TYPE_MAP.get(field_type):
                                fields.append((field_name, ts_type))

        if fields:
            self.models[node.name] = fields


def parse_model_file(path: str) -> dict[str, list[tuple[str, str]]]:
    """Parse a Django models.py file and find model definitions.

    Args:
        path(str): The path to a Django models.py file.

    Returns:
        A dictionary where keys are model class names and values are lists of
        tuples, each containing a field name and its corresponding TypeScript type.
    """
    with open(path) as file:
        tree = ast.parse(file.read())

    visitor = ModelVisitor()
    visitor.visit(tree)

    return visitor.models


def write_ts_type(_models: dict[str, list[tuple[str, str]]], output_path: str) -> None:
    """
    Write TypeScript type definitions for Django models to a file.

    Args:
        _models: A dictionary where keys are model class names and values are lists of
            tuples, each containing a field name and its corresponding TypeScript type.
        output_path: The path to the output file.
    """
    with open(output_path, "w") as file:
        for model_name, fields in _models.items():
            file.write(f"export type {model_name} = {{\n")
            for field_name, ts_type in fields:
                file.write(f"  {field_name}: {ts_type};\n")
            file.write("};\n\n")


def main(ignore_dirs: list[str]) -> None:
    """
    Main function to traverse directories, parse model files, and write TypeScript types.
    """
    # Create ts_models directory in the root directory
    os.makedirs("ts_models", exist_ok=True)

    # Recursively search for all models.py files
    with Progress() as progress:
        task = progress.add_task("[cyan]Starting...", total=100)

        for root, dirs, files in os.walk(".."):
            dirs[:] = [d for d in dirs if d not in ignore_dirs]

            for name in files:
                if name == "models.py":
                    file_path = os.path.join(root, name)

                    # Parse the Django model file
                    models = parse_model_file(file_path)

                    # Name the output file based on the directory where models.py was found
                    dir_name = os.path.basename(root)
                    output_file = os.path.join("ts_models", f"{dir_name}-models.ts")

                    # Write the TypeScript type to the output file
                    write_ts_type(models, output_file)

                    for model_name in models.keys():
                        progress.update(
                            task,
                            advance=10,
                            description=f"[cyan]Generating type for {model_name}...",
                        )
        progress.stop()
        rich.print("[green]Done!")


if __name__ == "__main__":
    parser = argparse.ArgumentParser(
        description="Convert Django models to TypeScript types."
    )
    parser.add_argument("--ignore_dirs", nargs="+", help="directories to ignore")
    args = parser.parse_args()
    main(args.ignore_dirs or [])