const { Router } = require("express");
const RegRouter = require("./register");
const PhotoRouter = require("./photo");
const AlbumRouter = require("./album");

const MainRouter = Router();

MainRouter.use("/api", RegRouter);
MainRouter.use("/api", PhotoRouter);
MainRouter.use("/api", AlbumRouter);

module.exports = MainRouter;
