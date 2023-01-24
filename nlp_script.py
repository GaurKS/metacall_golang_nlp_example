from nltk.corpus import stopwords
from nltk.tokenize import word_tokenize

def remove_stopwords(text):
    # Tokenize the text
    words = word_tokenize(text)
    # Get a set of stop words
    stop_words = set(stopwords.words("english"))
    # Remove stop words from the tokenized words
    filtered_words = [word for word in words if word.lower() not in stop_words]
    # Join the filtered words back into a single string
    filtered_text = " ".join(filtered_words)
    return filtered_text

# text = "This is an example of removing stop words from a string."
# print(remove_stopwords(text))
