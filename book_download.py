import os
import time
import requests
from bs4 import BeautifulSoup
from tqdm import tqdm

# Diretório onde os arquivos serão salvos
DOWNLOAD_DIR = "data/documents"
os.makedirs(DOWNLOAD_DIR, exist_ok=True)

# Base de busca
BASE_URL = "https://www.gutenberg.org"

# Pega os IDs dos primeiros 100 livros do catálogo
def get_top_100_book_ids():
    url = f"{BASE_URL}/ebooks/search/?sort_order=downloads"
    book_ids = set()
    page = 1

    while len(book_ids) < 100:
        resp = requests.get(url + f"&start_index={(page-1)*25+1}")
        soup = BeautifulSoup(resp.text, 'html.parser')
        links = soup.select('li.booklink a.link')

        for link in links:
            href = link.get("href")
            if href and href.startswith("/ebooks/"):
                book_id = href.split("/")[2]
                book_ids.add(book_id)
                if len(book_ids) >= 100:
                    break
        page += 1

    return list(book_ids)

# Faz o download de um livro em PDF, EPUB e TXT se disponíveis
def download_book(book_id):
    formats = {
        # "pdf": f"{BASE_URL}/ebooks/{book_id}.pdf.images", # Formatos PDFs não estão sendo encontrados
        "epub": f"{BASE_URL}/ebooks/{book_id}.epub.noimages",
        "txt": f"{BASE_URL}/files/{book_id}/{book_id}-0.txt"
    }

    for fmt, url in formats.items():
        try:
            response = requests.get(url, stream=True)
            if response.status_code == 200 and 'text/html' not in response.headers.get('Content-Type', ''):
                filename = os.path.join(f"{DOWNLOAD_DIR}/{fmt}", f"{book_id}.{fmt}")
                with open(filename, "wb") as f:
                    for chunk in response.iter_content(1024):
                        f.write(chunk)
                print(f"[✓] Downloaded: {filename}")
            else:
                raise Exception("Formato não disponível")
        except Exception as e:
            print(f"[x] {fmt.upper()} não disponível para o livro {book_id}.")

def main():
    print("Buscando os 100 livros mais populares do Project Gutenberg...")
    book_ids = get_top_100_book_ids()
    print(f"Total de livros coletados: {len(book_ids)}\n")

    for book_id in tqdm(book_ids, desc="Baixando livros"):
        download_book(book_id)
        time.sleep(1)  # Delay para evitar bloqueios

if __name__ == "__main__":
    main()
