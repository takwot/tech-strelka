import { useState } from "react";
import styles from "./Main.module.scss";
import Switch from "./Switch/Switch";
import Albums from "./Albums/Albums";

const Main = () => {
  const [album, setAlbum] = useState(true);
  return (
    <div className={styles.container}>
      <div className={styles.content}>
        <Switch album={album} setAlbum={setAlbum} />
        {album ? (
          <Albums />
        ) : (
          <div className={styles.photo_cont}>
            <div />
            <div />
            <div />
            <div />
            <div />
            <div />
            <div />
            <div />
            <div />
            <div />
            <div />
            <div />
            <div />
            <div />
            <div />
            <div />
            <div />
          </div>
        )}
      </div>
    </div>
  );
};

export default Main;
