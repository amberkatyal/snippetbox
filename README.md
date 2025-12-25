# snippetbox
Snippetbox lets people paste and share snippets of text -- a bit like Pastebin.

## Project Structure
**cmd** - cmd directory will contain the application specific code for the executable applications. Here its a web app which lives under **cmd/web**.

**internal** - contain the ancillary non-app specific code in the project. We'll use it to hold the potential reuseable code like validation helpers & SQL database models for the project.

**ui** - will contain the assets used by the web app. The ui/html will contain the html templates and ui/static will contain static(cdn) files like css and images.
