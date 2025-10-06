import express from "express";
import cors from "cors";
import helmet from "helmet";
import mailer from "./utils/Route.js";

const app = express();

app.use(express.json());
app.use(
  cors({
    origin: process.env.CORS_ORIGIN,
    methods: ["GET", "POST", "PUT", "DELETE"],
    allowedHeaders: ["Content-Type", "Authorization"],
  })
);
app.use(helmet());

const PORT = process.env.PORT || 3000;

app.use("/", (req, res) => {
  res.send("Unsafemail TS is running");
});

app.use("/api", mailer);

app.listen(PORT, () => {
  console.log(`Server is running on port ${PORT}`);
});
