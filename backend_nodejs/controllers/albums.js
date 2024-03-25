const pool = require("../db/db");

const createAlbum = (req, res) => {
  const { name, author, photos } = req.body;

  pool.query(
    "INSERT INTO albums (name, author, photos) VALUES ($1, $2, $3)",
    [name, author, JSON.stringify(photos)],
    (err, result) => {
      if (err) {
        console.log(err);
        res.status(500).json({ error: "Something went wrong" });
        return;
      } else {
        res.status(200).json({ message: "Album created successfully" });
        return;
      }
    }
  );
};

const getAllAlbums = (req, res) => {
  pool.query(
    "SELECT * FROM albums WHERE author = $1",
    [req.query.author],
    (err, result) => {
      if (err) {
        console.log(err);
        res.status(500).json({ error: "Something went wrong" });
        return;
      } else {
        res.status(200).json(result.rows);
        return;
      }
    }
  );
};

const deleteAlbum = (req, res) => {
  const { id } = req.query;

  pool.query("DELETE FROM albums WHERE id = $1", [id], (err, result) => {
    if (err) {
      console.log(err);
      res.status(500).json({ error: "Something went wrong" });
      return;
    } else {
      res.status(200).json({ message: "Album deleted successfully" });
      return;
    }
  });
};

module.exports = {
  createAlbum,
  getAllAlbums,
  deleteAlbum,
};
