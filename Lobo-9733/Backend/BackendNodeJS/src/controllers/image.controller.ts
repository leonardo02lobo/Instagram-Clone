import { Request, Response } from 'express';
import fs from 'fs';
import path from 'path';

const uploadDir = path.resolve(__dirname, '../../uploads');

export const uploadImage = (req: Request, res: Response) => {
  if (!req.file) return res.status(400).json({ message: 'No file uploaded' });

  res.json({
    message: 'File uploaded successfully',
    filename: req.file.filename,
    url: `/images/${req.file.filename}`
  });
};

export const getImage = (req: Request, res: Response) => {
  const filename = req.params.filename;
  const filePath = path.join(uploadDir, filename);

  if (!fs.existsSync(filePath)) {
    return res.status(404).json({ message: 'Image not found' });
  }

  res.sendFile(filePath);
};

export const listImages = (_req: Request, res: Response) => {
  const files = fs.readdirSync(uploadDir);
  res.json({ files });
};
