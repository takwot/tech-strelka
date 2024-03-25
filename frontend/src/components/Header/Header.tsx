import { NavLink, Outlet } from "react-router-dom";
import styles from "./Header.module.scss";
import { useEffect, useState } from "react";
import InsertPhotoIcon from "@mui/icons-material/InsertPhoto";
import useAuth from "../../core/store/auth";
import { useCookies } from "react-cookie";
import ExitToAppIcon from "@mui/icons-material/ExitToApp";
import Api from "../../core/api/api";

const Header = () => {
  const [menu, setMenu] = useState(false);

  const auth = useAuth((state) => state.isAuthenticated);
  const { user, setIsAuthenticated, setUser } = useAuth();

  const [cookies, setCookie, removeCookie] = useCookies();

  useEffect(() => {
    const name = cookies.name;
    const password = cookies.password;

    Api.login(name, password).then((res) => {
      if (res.data.user) {
        cookies.token;
        setIsAuthenticated(true);
        setCookie("asd", "asd");
        setUser(res.data.user);
      }
    });
  }, []);

  return (
    <>
      <div className={styles.container}>
        <NavLink to={"/"} className={styles.name}>
          <InsertPhotoIcon /> 300kPhotos
        </NavLink>
        {!auth && (
          <div style={{ display: "flex", alignItems: "center", gap: "10px" }}>
            <NavLink to={"/auth"}>
              <button>Register</button>
              <button>Login</button>
            </NavLink>
          </div>
        )}
        {auth && (
          <div style={{ display: "flex", alignItems: "center", gap: "10px" }}>
            <p onClick={() => setMenu(!menu)}>{user.name}</p>
            {/* <div className={styles.avatar} onClick={() => setMenu(!menu)}></div> */}
            <div
              onClick={() => {
                removeCookie("name");
                removeCookie("password");
                removeCookie("asd");
                window.location.reload();
              }}
            >
              <ExitToAppIcon sx={{ color: "white" }} />
            </div>
          </div>
        )}
      </div>
    </>
  );
};

export default Header;
