/** @type {import("tailwindcss").Config} */
// has to be migrated to: https://tailwindcss.com/docs/functions-and-directives#compatibility
module.exports = {
	content: ["renderer/page/**/*.html.tmpl"],
	theme: {
		fontFamily: {
			geometric: [
				"\"Euclid Circular B\", Futura, \"Trebuchet MS\", Arial, sans-serif",
				{ fontFeatureSettings: "\"kern\"" },
			],
			sans: [
				"\"Suisse Int'l\", -apple-system, BlinkMacSystemFont, Helvetica, Arial, sans-serif",
				{ fontFeatureSettings: "\"kern\"" },
			],
		},
		extend: {
			spacing: {
				128: "32rem",
				172: "48rem",
				256: "64rem",
			},
			fontSize: {
				"2xs": [
					"0.625rem",
					{
						lineHeight: "0.75rem",
						letterSpacing: "0.05rem",
					},
				],
			},
			transitionProperty: {
				right: "right",
			},
		},
	},
	plugins: [require("@tailwindcss/typography")],
};
