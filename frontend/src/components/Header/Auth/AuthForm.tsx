import { Button, TextField } from "@mui/material";
import styles from "./Auth.module.scss";
import { useState } from "react";
import Api from "../../../core/api/api";

const AuthForm = () => {
  const [register, setRegister] = useState(false);

  const [password, setPassword] = useState("");
  const [nick, setNick] = useState("");

  const auth = () => {
    Api.login(nick, password)
      .then((res) => {
        console.log(res.data);
      })
      .catch((err) => console.log(err.response.data));
  };

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
          label="Nickname"
          sx={{ fontSize: "14px", width: "100%" }}
          variant="standard"
          value={nick}
          onChange={(e) => setNick(e.target.value)}
        />
        {register && (
          <>
            <TextField
              label="Name"
              sx={{ fontSize: "14px", width: "100%" }}
              variant="standard"
            />
            <TextField
              label="Email"
              sx={{ fontSize: "14px", width: "100%" }}
              variant="standard"
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
          onClick={register ? auth : auth}
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
