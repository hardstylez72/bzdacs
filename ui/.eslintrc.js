module.exports = {
  root: true,
  env: {
    node: true,
  },
  extends: [
    // 'plugin:vue/recommended',
    '@vue/airbnb',
    '@vue/typescript/recommended',
  ],
  parserOptions: {
    ecmaVersion: 2020,
  },
  overrides: [
    {
      files: ['*.vue'],
      rules: {
        indent: 'off',
      },
    },
  ],
  rules: {
    'no-console': process.env.NODE_ENV === 'production' ? 'warn' : 'off',
    'no-debugger': process.env.NODE_ENV === 'production' ? 'warn' : 'off',
    'max-len': 'warn',
    'consistent-return': 'off',
    'class-methods-use-this': ['warn'],
    '@typescript-eslint/ban-ts-ignore': 0,
    'import/extensions': 0,
    'no-param-reassign': 0,
    '@typescript-eslint/camelcase': 0,
  },
};
