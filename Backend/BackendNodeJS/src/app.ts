import express from 'express';
import imageRoutes from './routes/image.router';
import cors from "cors"

const app = express();
    const corsOptions: cors.CorsOptions = {
      origin: ['http://localhost:4321'], 
      methods: 'GET,HEAD,PUT,PATCH,POST,DELETE', 
      credentials: true, 
      optionsSuccessStatus: 204, 
    };

    app.use(cors(corsOptions));

app.use(express.json());
app.use('/images', imageRoutes);

export default app;
