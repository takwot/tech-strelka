const Pool = require("pg").Pool;
const pool = new Pool({
  user: "dev",
  host: "45.84.225.194",
  database: "techstrelka",
  password: "f7636c0azc8bfk1",
  port: 5555,
});

module.exports = pool;
