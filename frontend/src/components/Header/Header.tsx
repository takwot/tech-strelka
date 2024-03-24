import { NavLink, Outlet } from "react-router-dom";
import styles from "./Header.module.scss";
import { useState } from "react";
import InsertPhotoIcon from "@mui/icons-material/InsertPhoto";
import useAuth from "../../core/store/auth";

const Header = () => {
  const [menu, setMenu] = useState(false);

  const auth = useAuth((state) => state.isAuthenticated);

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
            <p onClick={() => setMenu(!menu)}>NICKNAME</p>
            <div className={styles.avatar} onClick={() => setMenu(!menu)}></div>
          </div>
        )}
      </div>
    </>
  );
};

export default Header;
