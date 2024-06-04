import random

def get_word():
  """Gets a random word from a list."""
  words = ["python", "javascript", "programming", "computer", "science"]
  return random.choice(words)

def play_hangman():
  """Plays the hangman game."""
  word = get_word()
  word_letters = set(word)  # Letters in the word
  alphabet = set(chr(x) for x in range(97, 123))  # All lowercase letters
  used_letters = set()  # What the user has guessed

  lives = 6

  # Keep getting user input until the game is over
  while len(word_letters) > 0 and lives > 0:
    # Letters used
    print('You have used these letters: ', ' '.join(used_letters))

    # Current word (e.g. W - R D)
    word_list = [letter if letter in used_letters else '-'
 for letter in word]
    print('Current word: ', ' '.join(word_list))

    user_letter = input('Guess a letter: ').lower()
    if user_letter in alphabet - used_letters:
      used_letters.add(user_letter)
      if user_letter in word_letters:
        word_letters.remove(user_letter)
        print('')

      else:
        lives -= 1  # Takes away a life if wrong
        print('Letter is not in word.')

    elif user_letter in used_letters:
      print('You have already used that character. Please try again.')

    else:
      print('Invalid character. Please try again.')

  # Gets here when len(word_letters) == 0 OR lives == 0
  if lives == 0:
    print('You died, sorry. The word was', word)
  else:
    print('You guessed the word', word, '!!')

play_hangman()