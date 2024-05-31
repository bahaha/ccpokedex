# Pokedex

### Requirements Exploration

- [ ] Display basic information about each Pokemon (e.g., name, image)

- [ ] Allow user to click on a Pokemon to view detailed information in a modal or a new page.

- [ ] Enable filtering and searching of Pokemon by various attributes like type, abilities, …etc.

- [ ] Fetch data from the PokeAPI and render it efficiently in a user-friendly interface

- [ ] mobile-first design to focus on the functionalities of application

- [ ] Infinite scroll to get more pokemon to display with the same filters

### Non Functional Requirement

- [ ] view detailed information in a modal or a new page with [Shallow Routing](https://kit.svelte.dev/docs/shallow-routing)

---

### Architecture

- [ ] Go + Templ + HTMX

- [ ] Go (w/ API) + Sveltekit

- [ ] Go (w/ API) + CSR

---

### Data Model

- Pokemon: Represents each individual Pokemon. Fields include:

  - `id`: int

  - `name`: string

  - `type`: string[] (w/ enum ?)

  - `abilities`: string[]

  - `image_url`: string

- User Preferences (if implemented user account):

  - `user_id`: int

  - `favorite_pokemon`: int[]

---

### API

- `GET /api/pokemon`: Fetches a list of Pokemon, supports filters.

- `GET /api/pokemon/{id}`: Fetches detailed information about a specific Pokemon.

- `GET /web/pokemon`: Fetches a list of Pokemon, supports filters, return as HTML markup.

- `GET /web/pokemon/{id}`

- `POST /web/{user_id}/pokemon/favorite`: add pokemon to the favorite list

…

---

### Optimization

- Caching

- Lazy Loading

- Virtualisation

- Server-Send Event on top of HTTP/2

