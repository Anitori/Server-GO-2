const myInput = document.getElementById("myFile");
const valorMax = 25000000;


//Validación de tamaño de el archivo 

myInput.addEventListener("change", function () {
	// si no hay archivos, regresamos
	if (this.files.length <= 0) return;

	// Valido el primer archivo únicamente
	const archivo = this.files[0];

	if (archivo.size > valorMax) {
		const tamanioEnMb = valorMax / 1000000;
		alert(`El tamaño máximo es ${tamanioEnMb} MB`);
		// Limpiar
		myInput.value = "";
	}
});

//






