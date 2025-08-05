import multer from 'multer';
import path from 'path';
import fs from 'fs';

const uploadDir = path.resolve(__dirname, '../../uploads');

if (!fs.existsSync(uploadDir)) {
  fs.mkdirSync(uploadDir);
}

const storage = multer.diskStorage({
  destination: (_req, _file, cb) => {
    cb(null, uploadDir);
  },
  filename: (_req, file, cb) => {
    const ext = path.extname(file.originalname);
    const name = `${Date.now()}-${Math.random().toString(36).substr(2, 9)}${ext}`;
    cb(null, name);
  }
});

export const upload = multer({ storage });
