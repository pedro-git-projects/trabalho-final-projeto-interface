<template>
<div class="wrapper d-flex flex-column min-vh-100">
	<div class="container">
		<div class="row text-center">
			<div class="col mt-5">
				<h1>Cthullu Online</h1>
			</div>
		</div>
	</div>
	<div class="container mt-3">
		<div class="row text-center">
			<div class="col">
				<p>
					Bem vindo ao <b>Cthullu Online</b>, um site onde você
					pode gerar investigadores para a sétima edição do jogo 
					<i>Call of Cthullu</i> da Chaosium.
				</p>
<router-link :to="{ path: '/criar' }"><button type="button" class="btn btn-primary">Criar Agora</button></router-link>			
</div>
		</div>
	</div>
	<div class="container">
		<div class="row text-center">
			<div class="col">
				<img alt="Cthullu and Lovecraft" src="../assets/home_image_3.png" class="img-fluid">
			</div>
		</div>
	</div>
</div>
</template>

<style scoped>
	h1 {
		font-family: 'Macondo', cursive;
	}
</style>
