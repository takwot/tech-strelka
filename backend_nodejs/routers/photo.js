const { Router } = require("express");
const PhotoRouter = require("../controllers/photo");

const router = Router();

router.post("/photo", PhotoRouter.uploadPhoto);
router.get("/photo", PhotoRouter.getPhoto);
router.get("/photo/all", PhotoRouter.getAllUserPhoto);
router.delete("/photo", PhotoRouter.deletePhoto);

module.exports = router;
