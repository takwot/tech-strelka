import DeleteIcon from "@mui/icons-material/Delete";
import { useState } from "react";
import Api from "../../../core/api/api";

type Props = {
  src: string;
  getAll: () => void;
  isExist?: boolean;
};

const Image = ({ src, getAll, isExist }: Props) => {
  const [view, setView] = useState(false);

  const deleteHandler = () => {
    Api.deletePhoto(src).then(() => {
      getAll();
    });
  };

  const srcFile = `http://45.84.225.194:3000/api/photo?filename=${src}`;

  return (
    <div
      onMouseEnter={() => setView(true)}
      onMouseLeave={() => setView(false)}
      style={{
        display: "flex",
        flexDirection: "column",
        alignItems: "flex-end",
        gap: "5px",
      }}
    >
      <img src={srcFile} />
      {view && !isExist && (
        <div onClick={deleteHandler}>
          <DeleteIcon sx={{ color: "white", cursor: "pointer" }} />
        </div>
      )}
    </div>
  );
};

export default Image;
