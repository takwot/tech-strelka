import { Button } from "@mui/material";
import styles from "./Switch.module.scss";
import React, { useRef } from "react";
import Api from "../../../core/api/api";

type Props = {
  album: boolean;
  setAlbum: React.Dispatch<React.SetStateAction<boolean>>;
};

const Switch = ({ setAlbum, album }: Props) => {
  const ref: any = useRef();

  return (
    <div className={styles.switch}>
      <Button variant="contained" onClick={() => ref.current.click()}>
        Add
      </Button>
      <input
        ref={ref}
        hidden
        type="file"
        onChange={(e) => {
          const data = new FormData();

          const files: any = e.target.files;

          if (files) {
            data.append("file", files[0]);
            const token =
              "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTE0MDU2OTgsImlhdCI6MTcxMTM2MjQ5OCwidXNlcl9pZCI6Mn0.-nwDqnBIC6Qi9QYM-Va7NYAlZm07fSbLkezI-lwhm6U";
            Api.createPhoto(data, token).then((res) => {
              console.log(res.data);
            });
          }
        }}
      />
      <p onClick={() => setAlbum(!album)}>{album ? "All" : "Albums"}</p>
    </div>
  );
};

export default Switch;
