const pool = require("../db/db");
const fs = require("fs");

const uploadPhoto = (req, res) => {
  const file = req.files.file;
  const date = Date.now();
  const type = file.name.split(".").pop();
  const fileName = date + "." + type;
  file.mv("./files/" + fileName);

  const { id } = req.query;

  pool.query(
    "INSERT INTO photos (name, author) VALUES ($1, $2)",
    [fileName, id],
    (err, result) => {
      if (err) {
        console.log(err);
        res.status(500).json({ error: "Something went wrong" });
        return;
      } else {
        res.status(200).json({ message: "Photo uploaded successfully" });
        return;
      }
    }
  );
};

const getPhoto = (req, res) => {
  const { filename } = req.query;
  fs.readFile(`./files/${filename}`, (err, data) => {
    if (err) {
      console.log(err);
      res.json({ message: "Some error" }).status(400);
    }
    res.send(data);
  });
};

const getAllUserPhoto = (req, res) => {
  const { id } = req.query;
  pool.query("SELECT * FROM photos WHERE author = $1", [id], (err, result) => {
    if (err) {
      console.log(err);
      res.status(500).json({ error: "Something went wrong" });
      return;
    } else {
      res.status(200).json(result.rows);
      return;
    }
  });
};

const deletePhoto = (req, res) => {
  fs.unlink(`./files/${req.query.filename}`, (err) => {
    if (err) {
      console.log(err);
      res.status(500).json({ error: "Something went wrong" });
      return;
    }
    pool.query(
      "DELETE FROM photos WHERE name = $1",
      [req.query.filename],
      (err, result) => {
        if (err) {
          console.log(err);
          res.status(500).json({ error: "Something went wrong" });
          return;
        } else {
          res.status(200).json({ message: "Photo deleted successfully" });
          return;
        }
      }
    );
  });
};

module.exports = {
  uploadPhoto,
  getPhoto,
  getAllUserPhoto,
  deletePhoto,
};
