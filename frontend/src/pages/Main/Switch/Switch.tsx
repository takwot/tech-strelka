import { Button } from "@mui/material";
import styles from "./Switch.module.scss";
import React, { useRef } from "react";
import Api from "../../../core/api/api";
import useAuth from "../../../core/store/auth";
import AddAlbum from "./AddAlbum/AddAlbum";

type Props = {
  album: boolean;
  setAlbum: React.Dispatch<React.SetStateAction<boolean>>;
  getAll: () => void;
};

const Switch = ({ setAlbum, album, getAll }: Props) => {
  const ref: any = useRef();

  const id = useAuth((state) => state.user.id);
  const [view, setView] = React.useState(false);

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
          const files: any = e.target.files;
          const data = new FormData();
          console.log(files);
          data.append("file", files[0]);
          Api.createPhoto(data, id).then(() => {
            getAll();
          });
        }}
      />
      {view && <AddAlbum setView={setView} />}
      {album && (
        <Button
          variant="contained"
          style={{ width: "150px" }}
          onClick={() => setView(true)}
        >
          Add album
        </Button>
      )}
      <p onClick={() => setAlbum(!album)}>{album ? "All" : "Albums"}</p>
    </div>
  );
};

export default Switch;
