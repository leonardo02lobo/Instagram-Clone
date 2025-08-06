import { Router } from 'express';
import { upload } from '../utils/multer.config';
import { uploadImage, getImage, listImages } from '../controllers/image.controller';

const router = Router();

router.post('/upload', upload.single('image'), uploadImage);
router.get('/:filename', getImage);
router.get('/', listImages);

export default router;
