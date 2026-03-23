mans = ['srcom.1', 'srcom-inspect.1', 'srcom-trans.1']
if get_option('with_docs')
	a2x = find_program('asciidoctor')
	foreach m : mans
		custom_target(
			m,
			output: [m],
			input: [m + '.adoc'],
			command: [
				a2x,
				'-a', 'srcom-version=v' + meson.project_version(),
				'--backend', 'manpage',
				'@INPUT@',
				'-D', meson.current_build_dir(),
			],
			install: true,
			install_dir: join_paths(get_option('mandir'), 'man1'),
		)
		custom_target(
			m + '.html',
			output: [m + '.html'],
			input: [m + '.adoc'],
			command: [
				a2x,
				'-a', 'srcom-version=v' + meson.project_version(),
				'--backend', 'html',
				'@INPUT@',
				'-D', meson.current_build_dir(),
			],
			install_dir: get_option('datadir') / 'doc' / 'srcom',
		)
	endforeach
endif
