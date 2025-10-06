import type { Request, Response } from "express";
import { Router } from "express";
import { sendEmail } from "./Mailer.js";

const mailer = Router();

mailer.post("/send-email", async (req: Request, res: Response) => {
  const { to, subject, html } = req.body;

  try {
    const info = await sendEmail({ to, subject, html });
    res.status(200).json({ message: "Email sent successfully", info });
  } catch (error) {
    res.status(500).json({ message: "Error sending email", error });
  }
});

export default mailer;
