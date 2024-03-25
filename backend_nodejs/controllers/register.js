const pool = require("../db/db");

const registration = (req, res) => {
  const { name, password, email } = req.body;

  pool.query("SELECT * FROM users WHERE email = $1", [email], (err, result) => {
    if (err) {
      res.status(500).json({ error: "Something went wrong" });
      return;
    } else if (result.rowCount > 0) {
      res.status(400).json({ error: "User already exists" });
      return;
    }

    const password_hash = require("bcrypt").hashSync(password, 10);

    pool.query(
      "INSERT INTO users (name, password_hash, email) VALUES ($1, $2,$3)",
      [name, password_hash, email],
      (err, result) => {
        if (err) {
          console.log(err);
          res.status(500).json({ error: "Something went wrong" });
          return;
        } else {
          res.status(200).json({ message: "User created successfully" });
          return;
        }
      }
    );
    return;
  });
};

const login = (req, res) => {
  const { name, password } = req.body;
  pool.query("SELECT * FROM users WHERE name = $1", [name], (err, result) => {
    const user = result.rows[0];

    if (user) {
      const password_hash = require("bcrypt").compareSync(
        password,
        user.password_hash
      );

      console.log(password_hash);

      if (!password_hash) {
        res.status = 401;
        res.json({ error: "error" });
        return;
      } else {
        const { id, email, name, password_hash } = user;
        res.json({ user: { id, email, name } });
      }
    } else {
      res.json({ error: "User not found" });
    }
  });
};

module.exports = {
  registration,
  login,
};
