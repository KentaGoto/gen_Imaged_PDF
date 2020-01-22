# gen_Imaged_PDF

## Usage

1. **gen_Imaged_PDF.pl** をコマンドプロンプトから起動する。またはダブルクリックで起動する
2. PDF の入ったフォルダをドラッグアンドドロップする
3. 画像解像度を整数値で指定する
4. 処理完了後は PDF と同じ場所に同名のフォルダができ、その中に画像化された PDF と各ページの PNG が作成される。  
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
- Perl ( 5.26 or above) and Modules (PDF::API2::Lite, File::Find::Rule, File::Basename)
- mutool (* Put mutool.exe in the same directory as pdf2images_parallel.exe.)
