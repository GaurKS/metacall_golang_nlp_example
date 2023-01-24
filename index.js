const { remove_stopwords } = require('./nlp_script.py');

function main(text) {
  return remove_stopwords(text);
}

module.exports = {
  main
}