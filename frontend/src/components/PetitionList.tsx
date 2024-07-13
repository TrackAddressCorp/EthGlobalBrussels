import React, { useEffect, useState } from 'react';
import { List, ListItem, ListItemText, Paper, ButtonBase, Box, Typography } from '@mui/material';
import { ThemeProvider, createTheme, styled } from '@mui/material/styles';

interface PetitionData {
  ID: number;
  title: string;
  description: string;
  pdf_url: string | null;
  signs: number;
}

const theme = createTheme();

const StyledListItem = styled(ListItem)(({ theme }) => ({
  '&:hover': {
    backgroundColor: theme.palette.action.hover,
  },
  display: 'flex',
  justifyContent: 'space-between',
  alignItems: 'center',
}));

const SignCount = styled(Box)(({ theme }) => ({
  backgroundColor: theme.palette.grey[300],
  borderRadius: '50%',
  width: 40,
  height: 40,
  display: 'flex',
  justifyContent: 'center',
  alignItems: 'center',
}));

const App: React.FC = () => {
  const [petitions, setPetitions] = useState<PetitionData[]>([]);

  useEffect(() => {
    const fetchPetitions = async () => {
      try {
        const response = await fetch('http://localhost:4242/petitions');
        const data = await response.json();
        setPetitions(data.petitions);
      } catch (error) {
        console.error('Error fetching petitions:', error);
      }
    };

    fetchPetitions();
  }, []);

  return (
    <ThemeProvider theme={theme}>
      <Paper elevation={3} sx={{ maxWidth: 600, margin: '20px auto', padding: '20px' }}>
        <List>
          {petitions.map((petition) => (
            <ButtonBase
              key={petition.ID}
              component="a"
              href={petition.pdf_url || '#'}
              sx={{ width: '100%', textAlign: 'left' }}
            >
              <StyledListItem>
                <ListItemText
                  primary={petition.title}
                  secondary={petition.description}
                />
                <SignCount>
                  <Typography variant="body2" color="textSecondary">
                    {petition.signs}
                  </Typography>
                </SignCount>
              </StyledListItem>
            </ButtonBase>
          ))}
        </List>
      </Paper>
    </ThemeProvider>
  );
};

export default App;
