const express = require("express");
const albums = require("../controllers/albums");

const router = express.Router();

router.post("/album", albums.createAlbum);
router.get("/album", albums.getAllAlbums);
router.delete("/album", albums.deleteAlbum);

module.exports = router;
