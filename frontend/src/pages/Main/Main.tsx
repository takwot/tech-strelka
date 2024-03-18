import { NavLink } from "react-router-dom";
import styles from "./Main.module.scss";
import Post from "./Post/Post";

const Main = () => {
  return (
    <div className={styles.container}>
      <div className={styles.content}>
        <Post />
        <Post />
        <Post />
        <Post />
      </div>
    </div>
  );
};

export default Main;
