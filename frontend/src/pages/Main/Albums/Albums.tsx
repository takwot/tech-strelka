import { useEffect, useState } from "react";
import styles from "./Albums.module.scss";
import useAuth from "../../../core/store/auth";
import Api from "../../../core/api/api";
import DeleteIcon from "@mui/icons-material/Delete";
import useAlbums from "../../../core/store/album";

type Props = {};

const Albums = (props: Props) => {
  const [albums, setAlbums] = useState<any[]>([]);

  const id = useAuth((state) => state.user.id);

  const { create, setCreate } = useAlbums();

  const [photo, setPhoto] = useState<string[]>([]);

  useEffect(() => {
    Api.getAllAlbum(id).then((res) => {
      setAlbums(res.data);
      setCreate(false);
    });
  }, [create]);

  const [current, setCurrent] = useState(0);
  const [isWatching, setIsWatching] = useState(false);

  return (
    <div className={styles.container}>
      {isWatching ? (
        <>
          <div
            onClick={() => {
              setIsWatching(false);
            }}
          >
            Назад
          </div>
          <div className={styles.photo_cont}>
            {photo.map((el) => {
              const srcFile = `http://45.84.225.194:3000/api/photo?filename=${el}`;
              return <img src={srcFile} />;
            })}
          </div>
        </>
      ) : (
        <>
          {albums.map((el) => (
            <div
              style={{
                width: "100%",
                display: "flex",
                alignItems: "center",
                justifyContent: "space-between",
              }}
            >
              <p
                onClick={() => {
                  setCurrent(el.id);
                  setIsWatching(true);
                  setPhoto(JSON.parse(el.photos));
                }}
              >
                - {el.name}
              </p>
              <div
                onClick={() => {
                  Api.deleteAlbum(el.id).then(() => {
                    Api.getAllAlbum(id).then((res) => {
                      setAlbums(res.data);
                      setCreate(false);
                    });
                  });
                }}
              >
                <DeleteIcon sx={{ color: "white", cursor: "pointer" }} />
              </div>
            </div>
          ))}
        </>
      )}
    </div>
  );
};

export default Albums;
