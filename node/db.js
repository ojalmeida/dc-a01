const { Pool } = require('pg')
const config = require('./config')

const client = new Pool(config.database)
client.connect()

const QueryErr = 'error when trying to query data'


function getLastEntries(n) {

    let query = 'SELECT id, col_texto, col_dt FROM tb01 ORDER BY col_dt DESC LIMIT $1'
    let values = [n]

    return client.query(query, values)

}

module.exports.getLastEntries = getLastEntries
