import express from "express";
import dotenv from "dotenv";
import { prisma } from "./utils/prisma";

dotenv.config();

const app = express();

const PORT = 8080;

app.get("/", (req, res) => {

  res.json({ message: "hello" });
});

app.get("/users", async (req, res) => {
  const user = await prisma.user.findMany();
  res.json(user);
});

app.listen(PORT, () => {
  console.log(`App is listening on PORT ${PORT}`);
});
