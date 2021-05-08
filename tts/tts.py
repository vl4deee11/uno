from gtts import gTTS
import os
from argparse import ArgumentParser, Namespace


def parse_args():
    parser = ArgumentParser(description='Speech synthesis')
    parser.add_argument(
        '--text',
        type=str,
        required=True
    )
    return parser.parse_args()


if __name__ == '__main__':
    args = parse_args()
    gTTS(text=args.text, lang='en', slow=False).save("./gen.mp3")