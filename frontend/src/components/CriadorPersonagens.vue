<template>
<div class="wrapper d-flex flex-column min-vh-100">
	<div class="container">
		<div class="row">
			<div class="col">
				<h1 class="mt-5">Gerador de Investigadores</h1>
				<hr>
 <Form
      @submit="onSubmit"
      :validation-schema="schema"
      @invalid-submit="onInvalidSubmit"
    >
      <TextInput
        name="nome"
        type="text"
        label="Nome"
        placeholder="Nome do investigador"
      />
      <TextInput
        name="idade"
        type="number"
        label="Idade"
        placeholder="1-99"
      />
      <TextInput
        name="residencia"
        type="text"
        label="Residência"
        placeholder="Cidade de residência"
      />
	<TextInput
        name="nascimento"
        type="text"
        label="Local de nascimento"
        placeholder="Cidade natal do investigador"
      />      <TextInput
        name="ocupacao"
        type="text"
        label="Ocupação"
        placeholder="Profissão do investigador"
      />

      <button class="submit-btn" type="submit">Submit</button>
    </Form>			
	</div>
		</div>
	</div>
</div>
</template>


<script>
import { Form } from "vee-validate";
import * as Yup from "yup";
import TextInput from "./TextInput.vue";

export default {
  name: "App",
  components: {
    TextInput,
    Form,
  },
  setup() {
    function onSubmit(values) {
		const requestOptions = {
                method: "POST",
                body: JSON.stringify( {
					nome: values.nome,	
					idade: parseInt(values.idade),
					residencia: values.residencia,
					nascimento: values.nascimento,
					ocupacao: values.ocupacao,
					}
				),
            }			
        fetch("http://localhost:4000/v1/criar", requestOptions)
        .then((response) => response.json())
		.then((data) => {
			console.log(data)
		})			    
	}

    function onInvalidSubmit() {
      const submitBtn = document.querySelector(".submit-btn");
      submitBtn.classList.add("invalid");
      setTimeout(() => {
        submitBtn.classList.remove("invalid");
      }, 1000);
    }

    // Using yup to generate a validation schema
    // https://vee-validate.logaretm.com/v4/guide/validation#validation-schemas-with-yup
	let vazio = "Este campo é obrigatório"
	let intervalo = "Insira um número entre 1 e 99"
    const schema = Yup.object().shape({
	nome: Yup.string().required(vazio),
	idade: Yup.number(intervalo).max(99, intervalo).min(1,intervalo).required(vazio).typeError("Insira um número entre 1 e 99"),
	residencia: Yup.string().required(vazio),
	nascimento: Yup.string().required(vazio),
	ocupacao: Yup.string().required(vazio),
    });

    return {
      onSubmit,
      schema,
      onInvalidSubmit,
    };
  },
};
</script>

<style>
* {
  box-sizing: border-box;
}

:root {
  --primary-color: #0071fe;
  --error-color: #f23648;
  --error-bg-color: #fddfe2;
  --success-color: #21a67a;
  --success-bg-color: #e0eee4;
}

form {
  width: 300px;
  margin: 0px auto;
  padding-bottom: 60px;
}

.submit-btn {
  background: var(--primary-color);
  outline: none;
  border: none;
  color: #fff;
  font-size: 18px;
  padding: 10px 15px;
  display: block;
  width: 100%;
  border-radius: 7px;
  margin-top: 40px;
  transition: transform 0.3s ease-in-out;
  cursor: pointer;
}

.submit-btn.invalid {
  animation: shake 0.5s;
  /* When the animation is finished, start again */
  animation-iteration-count: infinite;
}

@keyframes shake {
  0% {
    transform: translate(1px, 1px);
  }
  10% {
    transform: translate(-1px, -2px);
  }
  20% {
    transform: translate(-3px, 0px);
  }
  30% {
    transform: translate(3px, 2px);
  }
  40% {
    transform: translate(1px, -1px);
  }
  50% {
    transform: translate(-1px, 2px);
  }
  60% {
    transform: translate(-3px, 1px);
  }
  70% {
    transform: translate(3px, 1px);
  }
  80% {
    transform: translate(-1px, -1px);
  }
  90% {
    transform: translate(1px, 2px);
  }
  100% {
    transform: translate(1px, -2px);
  }
}

.submit-btn:hover {
  transform: scale(1.1);
}
</style>

