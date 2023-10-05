// ***********************************************************
// This example support/component.ts is processed and
// loaded automatically before your test files.
//
// This is a great place to put global configuration and
// behavior that modifies Cypress.
//
// You can change the location of this file or turn off
// automatically serving support files with the
// 'supportFile' configuration option.
//
// You can read more here:
// https://on.cypress.io/configuration
// ***********************************************************

// Import commands.js using ES2015 syntax:
import "./commands.js";

// Alternatively you can use CommonJS syntax:
// require('./commands')

//#region General
import { mount } from "cypress/vue";
//import {} from "../../src/assets/base.css";
import "../../src/setup.js";
//#endregion
//#region WordsTable Component
//import "@formkit/auto-animate/vue";
import { i18n } from "../../src/locales/i18n.js";
//#endregion

// Augment the Cypress namespace to include type definitions for
// your custom command.
// Alternatively, can be defined in cypress/support/component.d.ts
// with a <reference path="./component" /> at the top of your spec.
declare global {
  namespace Cypress {
    interface Chainable {
      mount: typeof mount;
    }
  }
}
Cypress.Commands.add('mount', (component, options = {}) => {
  // Setup options object
  options.global = options.global || {}
  options.global.stubs = options.global.stubs || {}
  options.global.components = options.global.components || {}
  options.global.plugins = options.global.plugins || []

  // Use store passed in from options, or initialize a new one
  const {/* store = getStore(), */ ...mountOptions} = options

  // Add plugins here
  options.global.plugins.push(i18n)

  return mount(component, mountOptions)
})
//Cypress.Commands.add("mount", mount);
