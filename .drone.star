def main(ctx):
  t = testing(ctx)
  c = changes(ctx)
  n = notify(ctx)

  c['depends_on'].append(t['name'])
  n['depends_on'].append(t['name'])
  n['depends_on'].append(c['name'])

  return [ t, c, n ]

def testing(ctx):
  return {
    'kind': 'pipeline',
    'type': 'docker',
    'name': 'testing',
    'platform': {
      'os': 'linux',
      'arch': 'amd64',
    },
    'steps': [
      {
        'name': 'goenv',
        'image': 'webhippie/golang:1.16',
        'pull': 'always',
        'commands': [
          'go env',
        ],
        'volumes': [
          {
            'name': 'gopath',
            'path': '/srv/app',
          },
        ],
      },
      {
        'name': 'vet',
        'image': 'webhippie/golang:1.16',
        'pull': 'always',
        'commands': [
          'make vet',
        ],
        'volumes': [
          {
            'name': 'gopath',
            'path': '/srv/app',
          },
        ],
      },
      {
        'name': 'staticcheck',
        'image': 'webhippie/golang:1.16',
        'pull': 'always',
        'commands': [
          'make staticcheck',
        ],
        'volumes': [
          {
            'name': 'gopath',
            'path': '/srv/app',
          },
        ],
      },
      {
        'name': 'lint',
        'image': 'webhippie/golang:1.16',
        'pull': 'always',
        'commands': [
          'make lint',
        ],
        'volumes': [
          {
            'name': 'gopath',
            'path': '/srv/app',
          },
        ],
      },
      {
        'name': 'build',
        'image': 'webhippie/golang:1.16',
        'pull': 'always',
        'commands': [
          'make build',
        ],
        'volumes': [
          {
            'name': 'gopath',
            'path': '/srv/app',
          },
        ],
      },
      {
        'name': 'test',
        'image': 'webhippie/golang:1.16',
        'pull': 'always',
        'commands': [
          'make test',
        ],
        'volumes': [
          {
            'name': 'gopath',
            'path': '/srv/app',
          },
        ],
      },
      {
        'name': 'codacy',
        'image': 'dronehippie/codacy:1',
        'pull': 'always',
        'settings': {
          'token': {
            'from_secret': 'codacy_token',
          },
        },
      },
    ],
    'volumes': [
      {
        'name': 'gopath',
        'temp': {},
      },
    ],
    'trigger': {
      'ref': [
        'refs/heads/master',
        'refs/tags/**',
        'refs/pull/**',
      ],
    },
  }

def changes(ctx):
  return {
    'kind': 'pipeline',
    'type': 'docker',
    'name': 'changes',
    'platform': {
      'os': 'linux',
      'arch': 'amd64',
    },
    'clone': {
      'disable': True,
    },
    'steps': [
      {
        'name': 'clone',
        'image': 'dronehippie/git:1',
        'pull': 'always',
        'settings': {
          'actions': [
            'clone',
          ],
          'remote': 'https://github.com/%s' % (ctx.repo.slug),
          'branch': ctx.build.source if ctx.build.event == 'pull_request' else 'master',
          'path': '/drone/src',
          'netrc_machine': 'github.com',
          'netrc_username': {
            'from_secret': 'github_username',
          },
          'netrc_password': {
            'from_secret': 'github_token',
          },
        },
      },
      {
        'name': 'changes',
        'image': 'dronehippie/chglog:1',
        'pull': 'always',
        'settings': {
          'output': 'CHANGELOG.md',
          'print': True,
        },
      },
      {
        'name': 'upload',
        'image': 'dronehippie/git:1',
        'pull': 'always',
        'settings': {
          'actions': [
            'commit',
            'push',
          ],
          'message': 'Automated changelog update [skip ci]',
          'branch': 'master',
          'author_email': 'drone@webhippie.de',
          'author_name': 'Drone',
          'netrc_machine': 'github.com',
          'netrc_username': {
            'from_secret': 'github_username',
          },
          'netrc_password': {
            'from_secret': 'github_token',
          },
        },
        'when': {
          'ref': {
            'exclude': [
              'refs/pull/**',
            ],
          },
        },
      },
    ],
    'depends_on': [],
    'trigger': {
      'ref': [
        'refs/heads/master',
      ],
    },
  }

def notify(ctx):
  return {
    'kind': 'pipeline',
    'type': 'docker',
    'name': 'notify',
    'platform': {
      'os': 'linux',
      'arch': 'amd64',
    },
    'clone': {
      'disable': True,
    },
    'steps': [
      {
        'name': 'telegram',
        'image': 'dronehippie/telegram:1',
        'pull': 'always',
        'settings': {
          'token': {
            'from_secret': 'telegram_token',
          },
          'recipient': {
            'from_secret': 'telegram_recipient',
          },
        },
      },
    ],
    'depends_on': [],
    'trigger': {
      'ref': [
        'refs/heads/master',
        'refs/tags/**',
      ],
      'status': [
        'failure',
      ],
    },
  }
