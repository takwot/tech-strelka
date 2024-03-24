import { Button, TextField } from "@mui/material";
import styles from "./Auth.module.scss";
import { useEffect, useState } from "react";
import Api from "../../../core/api/api";
import { useCookies } from "react-cookie";
import useAuth from "../../../core/store/auth";

const AuthForm = () => {
  const [register, setRegister] = useState(false);

  const { setIsAuthenticated } = useAuth();

  const [password, setPassword] = useState("");
  const [nick, setNick] = useState("");
  const [email, setEmail] = useState("");

  const [cookies, setCookies] = useCookies();

  const auth = () => {
    Api.login(nick, password)
      .then((res) => {
        cookies.token;
        if (res.data.token) {
          setCookies("token", res.data.token);
          setIsAuthenticated(true);
        }
      })
      .catch((err) => console.log(err.response.data));
  };

  const registerHandler = () => {
    Api.register(nick, password, email)
      .then(() => {
        setRegister(false);
      })
      .catch((err) => console.log(err.response.data));
  };

  useEffect(() => {
    setEmail("");
    setNick("");
    setPassword("");
  }, [register]);

  return (
    <div className={styles.auth}>
      <div
        style={{
          display: "flex",
          flexDirection: "column",
          gap: "12px",
          width: "100%",
          alignItems: "center",
        }}
      >
        <p style={{ fontSize: "24px", fontWeight: "600" }}>
          {register ? "Register" : "Login"}
        </p>
        <TextField
          label="Name"
          sx={{ fontSize: "14px", width: "100%" }}
          variant="standard"
          value={nick}
          onChange={(e) => setNick(e.target.value)}
        />
        {register && (
          <>
            <TextField
              label="Email"
              sx={{ fontSize: "14px", width: "100%" }}
              variant="standard"
              onChange={(e) => setEmail(e.target.value)}
              value={email}
            />
          </>
        )}
        <TextField
          label="Password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          type="password"
          sx={{ fontSize: "14px", width: "100%" }}
          variant="standard"
        />
      </div>
      <div
        style={{
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
          gap: "12px",
          width: "100%",
        }}
      >
        <Button
          onClick={register ? registerHandler : auth}
          sx={{ width: "100%" }}
          variant="contained"
        >
          {register ? "Registration" : "Login"}
        </Button>
        <p style={{ cursor: "pointer" }} onClick={() => setRegister(!register)}>
          {register ? "Login" : "Register"}
        </p>
      </div>
    </div>
  );
};

export default AuthForm;
