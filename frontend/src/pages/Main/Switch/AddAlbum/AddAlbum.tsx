import { Button } from "@mui/material";
import styles from "./AddAlbum.module.scss";
import useAuth from "../../../../core/store/auth";
import React, { useEffect, useState } from "react";
import Api from "../../../../core/api/api";
import useAlbums from "../../../../core/store/album";

type Props = {
  setView: React.Dispatch<React.SetStateAction<boolean>>;
};

const AddAlbum = ({ setView }: Props) => {
  const id = useAuth((state) => state.user.id);

  const [photos, setPhotos] = useState<string[]>([]);
  const [allPhotos, setAllPhotos] = useState<any[]>([]);
  const [name, setName] = useState("");
  const { setCreate } = useAlbums();

  useEffect(() => {
    Api.getAllPhoto(id).then((res) => {
      setAllPhotos(res.data);
    });
  }, []);

  const addHandle = () => {
    Api.createAlbum(name, id, photos).then((res) => {
      setPhotos([]);
      setCreate(true);
      setName("");
      setView(false);
    });
  };

  return (
    <div className={styles.container} onClick={() => setView(false)}>
      <div
        className={styles.cont}
        onClick={(e) => {
          e.stopPropagation();
        }}
      >
        <input
          placeholder="Album name"
          value={name}
          onChange={(e) => setName(e.target.value)}
        />
        <div
          style={{
            display: "flex",
            flexWrap: "wrap",
            gap: "10px",
            justifyContent: "center",
          }}
        >
          {allPhotos.map((el) => {
            const name = el.name;
            const url = `http://localhost:3000/api/photo?filename=${el.name}`;
            return (
              <div
                className={styles.img}
                key={el.id}
                style={{ maxWidth: "120px" }}
                onClick={() => {
                  const isExist = photos.find((el) => el === name);
                  if (isExist) {
                    setPhotos(photos.filter((el) => el !== name));
                  } else {
                    setPhotos([...photos, name]);
                  }
                }}
              >
                <img
                  src={url}
                  alt=""
                  style={{
                    border: photos.find((elem) => elem === el.name)
                      ? "1px solid red"
                      : "",
                  }}
                />
              </div>
            );
          })}
        </div>
        <Button variant="contained" onClick={addHandle}>
          Add
        </Button>
      </div>
    </div>
  );
};

export default AddAlbum;
