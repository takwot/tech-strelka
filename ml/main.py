from imageai.Detection import ObjectDetection
from fastapi import FastAPI, UploadFile
from fastapi.middleware.cors import CORSMiddleware
from itertools import groupby
import os


app = FastAPI()
current_direction = os.getcwd()

origins = ["http://localhost:5173"]

app.add_middleware(
    CORSMiddleware,
    allow_origins=origins,
    allow_credentials=True,
    allow_methods=[" * "],
    allow_headers=[" * "],
)


@app.get("/")
async def read_root():
    return {"Hello": "World"}

@app.post('/api/detect')
async def detect_file(file: UploadFile):
    with open('new_image.png', 'wb') as file_obj :
        file_obj.write(file.file.read()) 


    detector = ObjectDetection()
    detector.setModelTypeAsYOLOv3()
    detector.setModelPath(os.path.join(current_direction, 'yolov3.pt'))
    detector.loadModel()
    detections = detector.detectObjectsFromImage(
            input_image=os.path.join(current_direction, 'new_image.png')
        )
    

    objectDetected = []
    for eachObject in detections:
        if eachObject['percentage_probability'] > 90:
            objectDetected.append('#' + eachObject['name'])

    os.remove(os.path.join(current_direction, 'new_image.png'))
    
    withOutRepeat = [el for el, _ in groupby(objectDetected)]
    
    return {"tags": withOutRepeat}
