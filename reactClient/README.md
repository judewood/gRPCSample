# 🚀 Welcome to your new awesome project!

This project has been created using **webpack-cli**, you can now run

```
npm run build
```

or

```
yarn build
```

to bundle your application


things I did to fixup 

1. Module not found: Error: Can't resolve 'stream' in '... - 

Answer: followed instrucitons at https://stackoverflow.com/questions/74462296/vue-error-running-using-node-rdkafka-package  - led to setup webpack with webpackconfig here - https://webpack.js.org/configuration/ - this created a webpack.config.js file - I added the stack overflow stream recommendation to it

then needed to install prettier - instrucitons here https://prettier.io/docs/en/install.html

This whole thing ended up adding index.js and index.css - gave me hello world app

tried updating webpack.config.js to change to entry: "./src/App.tsx", - 

it is now unhappy and complaining
Could not find a required file.
  Name: index.js

This linnk https://stackoverflow.com/questions/64720086/could-not-find-a-required-file-adding-typescript-to-react-project suggests updating the react-scripts but mine is OK






