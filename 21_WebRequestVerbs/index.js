import express from "express";

const app = express();
const port = 3000;

app.use(express.json())
app.use(express.urlencoded({ extended: true }))

app.use(express.static("public"));

app.get("/", (req, res) => {
    res.status(200).send("Welcome to LearnCodeOnnline");
});

app.get("/get", (req, res) => {
    res.status(200).send("Hello from LearnCodeOnline");
});

app.post("/post", (req, res) => {
    let myJson = req.body;
    res.status(200).send(myJson);
});

app.post("/postform", (req, res) => {
    res.status(200).send(JSON.stringify(req.body));
});

app.listen(port, () => {
    console.log(`Server running on port ${port}`);
});
