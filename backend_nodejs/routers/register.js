const { Router } = require("express");
const register = require("../controllers/register.js");

const router = Router();

router.post("/auth/register", register.registration);
router.post("/auth/login", register.login);

module.exports = router;
