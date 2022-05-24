<template>
<div class="wrapper d-flex flex-column min-vh-100">
	<div class="container">
		<div class="row">
			<div class="col">
				<h1 class="mt-5">Gerador de Investigadores</h1>
				<hr>
				<form method="post" action="/criar" class="needs-validation" novalidate> 
				<text-input
					label="Nome"
					type="text"
					name="nome"
					required="true"
				>
				</text-input>

				<text-input
					label="Idade"
					type="number"
					name="idade"
					required="true"
				>
				</text-input>

				<text-input
					label="Residência"
					type="text"
					name="residencia"
					required="true"
				>
				</text-input>

				<text-input
					label="Local de Nascimento"
					type="text"
					name="local de nascimento"
					required="true"
				>
				</text-input>

				<text-input
					label="Ocupação"
					type="text"
					name="ocupacao"
					required="true"
				>
				</text-input>

				<hr>
				
				<input type="submit" class="btn btn-primary" value="criar">
				</form>
			</div>
		</div>
	</div>
</div>
</template>

<script>
import TextInput from './TextInput.vue'
export default {
	name: 'criar-investigador',
	components: {
		TextInput,
	},
 mounted() {
        (function () {
            'use strict'

            // Fetch all the forms we want to apply custom Bootstrap validation styles to
            var forms = document.querySelectorAll('.needs-validation')

            // Loop over them and prevent submission
            Array.prototype.slice.call(forms)
                .forEach(function (form) {
                form.addEventListener('submit', function (event) {
                    if (!form.checkValidity()) {
                    event.preventDefault()
                    event.stopPropagation()
                    }

                    form.classList.add('was-validated')
                }, false)
                })
            })()
    }}
</script>
