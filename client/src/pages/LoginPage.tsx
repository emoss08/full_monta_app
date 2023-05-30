/*
 * COPYRIGHT(c) 2023 MONTA
 *
 * This file is part of Monta.
 *
 * The Monta software is licensed under the Business Source License 1.1. You are granted the right
 * to copy, modify, and redistribute the software, but only for non-production use or with a total
 * of less than three server instances. Starting from the Change Date (November 16, 2026), the
 * software will be made available under version 2 or later of the GNU General Public License.
 * If you use the software in violation of this license, your rights under the license will be
 * terminated automatically. The software is provided "as is," and the Licensor disclaims all
 * warranties and conditions. If you use this license's text or the "Business Source License" name
 * and trademark, you must comply with the Licensor's covenants, which include specifying the
 * Change License as the GPL Version 2.0 or a compatible license, specifying an Additional Use
 * Grant, and not modifying the license in any other way.
 */

import React from "react";
import { useNavigate } from "react-router-dom";
import { useAuthStore } from "@/stores/authStore";
import { useErrorStore } from "@/stores/errorStore";
import { Loader2 } from "lucide-react";
import { useUserStore } from "@/stores/userStore";
import {
  Anchor,
  Checkbox,
  Container,
  Group,
  Paper,
  PasswordInput,
  TextInput,
  Text,
  Title,
  Button,
} from "@mantine/core";
import { isNotEmpty, useForm } from "@mantine/form";
import { LoginFormValues } from "@/types/login";
import axios from "@/lib/axiosConfig";
import { useLocalStorage } from "@mantine/hooks";

const LoginPage: React.FC = () => {
  const [isAuthenticated, setIsAuthenticated] = useAuthStore((state) => [
    state.isAuthenticated,
    state.setIsAuthenticated,
  ]);
  const [, setUser] = useUserStore((state) => [state.user, state.setUser]);
  const [loading, setLoading] = React.useState<boolean>(false);
  const [, setUserInfo] = useLocalStorage({
    key: "mt_user_info",
    serialize: (value: any) => JSON.stringify(value, null, 2),
    deserialize: (localStorageValue) => JSON.parse(localStorageValue || "{}"),
  });
  const form = useForm<LoginFormValues>({
    initialValues: {
      username: "",
      password: "",
    },

    validate: {
      username: isNotEmpty("Username is required."),
      password: isNotEmpty("Password is required."),
    },
  });

  const navigate = useNavigate();
  React.useEffect((): void => {
    if (isAuthenticated) {
      navigate("/");
    }
  }, [isAuthenticated, navigate]);

  const login = async (values: LoginFormValues) => {
    setLoading(true);
    try {
      const response = await axios.post("login/", {
        username: values.username,
        password: values.password,
      });

      setUserInfo(response.data);
      // localStorage.setItem("mt_user_info", JSON.stringify(response.data));
      setIsAuthenticated(true);
      setUser(response.data);
    } catch (error: any) {
      if (error.response && error.response.status === 400) {
        const errors = error.response.data.errors;
        console.log(errors);
      }
    } finally {
      setLoading(false);
    }
  };

  return (
    <>
      <Container size={450} my={50}>
        <Title align="center">Welcome back!</Title>
        <Text color="dimmed" size="sm" align="center" mt={5}>
          Do not have an account yet?{" "}
          <Anchor size="sm" component="button">
            Create account
          </Anchor>
        </Text>

        <Paper withBorder shadow="md" p={30} mt={30} radius="md">
          <form
            onSubmit={form.onSubmit((values: LoginFormValues) => {
              login(values).then(() => {});
            })}
          >
            <TextInput
              label="Username"
              placeholder="Your Username"
              withAsterisk
              {...form.getInputProps("username")}
            />
            <PasswordInput
              label="Password"
              placeholder="Your password"
              mt="md"
              withAsterisk
              {...form.getInputProps("password")}
            />
            <Group position="apart" mt="lg">
              <Checkbox label="Remember me" />
              <Anchor component="button" size="sm">
                Forgot password?
              </Anchor>
            </Group>
            <Button type="submit" fullWidth mt="xl" loading={loading}>
              Sign in
            </Button>
          </form>
        </Paper>
      </Container>
    </>
  );
};
export default LoginPage;