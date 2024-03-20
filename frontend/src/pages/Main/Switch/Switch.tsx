import { Button } from "@mui/material";
import styles from "./Switch.module.scss";
import React from "react";

type Props = {
  album: boolean;
  setAlbum: React.Dispatch<React.SetStateAction<boolean>>;
};

const Switch = ({ setAlbum, album }: Props) => {
  return (
    <div className={styles.switch}>
      <Button variant="contained">Add</Button>
      <p onClick={() => setAlbum(!album)}>{album ? "All" : "Albums"}</p>
    </div>
  );
};

export default Switch;
