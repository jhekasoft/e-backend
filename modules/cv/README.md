# CV module

CV module contain CV-data in the YAML format (at `data/cv` directory) and have API HTTP-endpoints for [jhekasoft.github.io (efremov.dev)](https://github.com/jhekasoft/jhekasoft.github.io) site.

CV module also can generate LaTeX CV.

## Generate LaTeX CV

1. Get `LaTeX` body:

```
GET /cv/latex
```

http://localhost:1988/cv/latex

2. Paste boty to the `cv.tex`.

3. Run command:

```bash
pdflatex cv.tex
```

4. Open `cv.pdf`:

```bash
open cv.pdf
```
