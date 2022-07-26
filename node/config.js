const yaml = require('js-yaml');
const fs   = require('fs');
let config = {database: {}, api: {}}

const ReadErr = 'error when trying to read configuration file'

try {
    config = yaml.load(fs.readFileSync('/etc/config.yaml'), 'utf8');
} catch (e) {
    console.log(`${ReadErr} : ${e.message}`);
}

module.exports = config

