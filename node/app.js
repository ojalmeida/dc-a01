const express = require('express')
const storage = require('./db')
const config = require('./config')
const {json} = require("express");
const uuid = require('uuid')

const app = express()
const port = config.api.port
const hostname = config.api.hostname


app.use(function (req, res, next) {
    res.setHeader('Access-Control-Allow-Origin', '*');
    res.setHeader('Access-Control-Allow-Methods', '*');
    res.setHeader('Access-Control-Allow-Headers', '*');
    next();
});

app.get('/tb01', async (req, res) => {

    let init =  Date.now()
    let id = uuid.v4()

    storage.getLastEntries(10).then(result => {

        rows = result.rows

        let out = []
        rows.forEach(row => {

            let outRow = {id: undefined, text: undefined, date: undefined}
            let date = new Date(row.col_dt)

            outRow.date = `${date.getDay()}/${date.getMonth()}/${date.getFullYear()} ${date.getHours()}:${date.getMinutes()}:${date.getSeconds()}`
            outRow.id = row.id
            outRow.text = row.col_texto

            out.push(outRow)

        })

        console.log(`INFO (node-api): ${id} ${req.ip} ${req.method} ${req.url} ${req.header("referer")} ${req.header("user-agent")} ${200} ${(Date.now() - init)/1000}`)
        res.send(out)


    }).catch(err => {

        res.status(500).send("An error ocurred")
        console.error(`ERROR (node-api): ${id} ${req.ip} ${req.method} ${req.url} ${req.header("referer")} ${req.header("user-agent")} ${500} ${(Date.now() - init) / 1000} | ${err.message}`)

    })


})

app.listen(port, hostname,() => {

    console.log(`API listening on ${hostname}:${port}`)

})

