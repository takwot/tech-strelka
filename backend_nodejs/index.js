const express = require("express");
const MainRouter = require("./routers/index");
const files = require("express-fileupload");
const cors = require("cors");
const pool = require("./db/db");

const app = express();

app.use(cors({ credentials: false, origin: "*" }));
app.use(express.json({}));
app.use(files({}));
app.use(express.urlencoded({ extended: true }));
app.use(MainRouter);

const port = 3000;

app.listen(port, () => {
  console.log("Server is running on port 3000");
});
