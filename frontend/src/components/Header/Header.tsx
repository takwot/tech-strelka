import { NavLink, Outlet } from "react-router-dom";
import styles from "./Header.module.scss";
import { useState } from "react";
import Person2Icon from "@mui/icons-material/Person2";
import PeopleAltIcon from "@mui/icons-material/PeopleAlt";
import PhotoLibraryIcon from "@mui/icons-material/PhotoLibrary";
import ExitToAppIcon from "@mui/icons-material/ExitToApp";
import InsertPhotoIcon from "@mui/icons-material/InsertPhoto";

const Header = () => {
  const [menu, setMenu] = useState(false);
  const [auth, setAuth] = useState(false);

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
      {auth && menu && (
        <div className={styles.menu}>
          <div className={styles.content}>
            <NavLink onClick={() => setMenu(false)} to={"/profile"}>
              <Person2Icon sx={{ color: "black" }} /> Профиль
            </NavLink>
            <NavLink onClick={() => setMenu(false)} to={"/profile"}>
              <PeopleAltIcon sx={{ color: "black" }} />
              Друзья
            </NavLink>
            <NavLink onClick={() => setMenu(false)} to={"/profile"}>
              <PhotoLibraryIcon sx={{ color: "black" }} />
              Альбомы
            </NavLink>
            <NavLink onClick={() => setMenu(false)} to={"/profile"}>
              <ExitToAppIcon sx={{ color: "black" }} /> Выйти
            </NavLink>
          </div>
        </div>
      )}
      <Outlet />
    </>
  );
};

export default Header;
