<template>
<div class="wrapper d-flex flex-column min-vh-100">
	<div class="container">
		<div class="row">
			<div class="col">
				<h1 class="mt-5">Gerador de Investigadores</h1>
				<hr>
                <form-tag @myevent="submitHandler" name="myform" event="myevent">
				<text-input
					:value="nome"
					@input="event => nome = event.target.value"
					label="Nome"
					type="text"
					name="nome"
					required="true">
				</text-input>

				<text-input
					:value="idade"
					@input="event => idade = event.target.value"
					label="Idade"
					type="number"
					name="idade"
					required="true"
				>
				</text-input>

				<text-input
					:value="residencia"
					@input="event => residencia = event.target.value"
					label="Residência"
					type="text"
					name="residencia"
					required="true"
				>
				</text-input>

				<text-input
					:value="nascimento"
					@input="event => nascimento = event.target.value"
					label="Local de Nascimento"
					type="text"
					name="nascimento"
					required="true"
				>
				</text-input>

				<text-input
					:value="ocupacao"
					@input="event => ocupacao = event.target.value"
					label="Ocupação"
					type="text"
					name="ocupacao"
					required="true"
				>
				</text-input>

				<hr>
				
				<input type="submit" class="btn btn-primary" value="Criar">
				</form-tag>
			</div>
		</div>
	</div>
</div>
</template>

<script>
import TextInput from './TextInput.vue'
import FormTag from './FormTag.vue'
export default {
	name: 'criar-investigador',
	components: {
		FormTag,
		TextInput,
	},
	data() {
		return {
			nome: null, 
			idade:"", 
			residencia: "", 
			nascimento: "", 
			ocupacao: "",
		}
	},
	methods: {
		submitHandler() {
		const requestOptions = {
                method: "POST",
                body: JSON.stringify( {
					nome: this.nome,	
					idade: parseInt(this.idade),
					residencia: this.residencia,
					nascimento: this.nascimento,
					ocupacao: this.ocupacao,
					}
				),
            }			

        fetch("http://localhost:4000/v1/criar", requestOptions)
        .then((response) => response.json())
		.then((data) => {
			console.log(data)
		})

		}
	},
}
</script>
