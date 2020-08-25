# gen_Imaged_PDF

## Description  

Convert each PDF page to PNG.  

## Usage

1. Start **gen_Imaged_PDF.pl** from the command prompt. Or start by double click
2. Drag and drop the folder containing the PDF
3. Specify the image resolution with an integer value
4. After processing is completed, a folder with the same name is created in the same location as the PDF, and an imaged PDF and a PNG of each page are created in that folder.
   **Example:**

```
Dir: D:\tool\Perl_Test\gen_Imaged_PDF\tryme
Resolution: 72

Processing...
Output PDF: D:\tool\Perl_Test\gen_Imaged_PDF\tryme\sample\sample.pdf
Output PDF: D:\tool\Perl_Test\gen_Imaged_PDF\tryme\ああああああああ\Translation_Memory_Management_QSG_en\Translation_Memory_Management_QSG_en.pdf

Done!
```

## Requires

- Windows
- Perl5
- mutool (* Put mutool.exe in the same directory as pdf2images_parallel.exe.)
