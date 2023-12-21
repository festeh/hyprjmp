/** @type {import('tailwindcss').Config}*/
import { join } from 'path';

// 1. Import the Skeleton plugin
import { skeleton } from '@skeletonlabs/tw-plugin';

const config = {
  content: ['./src/**/*.{html,js,svelte,ts}',
    join(require.resolve(
      '@skeletonlabs/skeleton'),
      '../**/*.{html,js,svelte,ts}'
    )
  ],

  theme: {
    extend: {}
  },

  plugins: [
    skeleton({
      themes: { preset: [ "skeleton" ] }
    })
  ]
};

module.exports = config;
