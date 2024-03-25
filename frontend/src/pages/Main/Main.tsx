import { useEffect, useState } from "react";
import styles from "./Main.module.scss";
import Switch from "./Switch/Switch";
import Albums from "./Albums/Albums";
import Api from "../../core/api/api";
import useAuth from "../../core/store/auth";
import Image from "./Image/Image";

const Main = () => {
  const [album, setAlbum] = useState(false);

  const [allPhotos, setAllPhotos] = useState<any[]>([]);

  const id = useAuth((state) => state.user.id);

  const getAll = () => {
    Api.getAllPhoto(id).then((res) => {
      console.log(res.data);
      setAllPhotos(res.data);
    });
  };

  useEffect(() => {
    getAll();
  }, []);

  return (
    <div className={styles.container}>
      <div className={styles.content}>
        <Switch getAll={getAll} album={album} setAlbum={setAlbum} />
        {album ? (
          <Albums />
        ) : (
          <div className={styles.photo_cont}>
            {allPhotos.map((el) => {
              return <Image isExist={false} getAll={getAll} src={el.name} />;
            })}
          </div>
        )}
      </div>
    </div>
  );
};

export default Main;
