// PetitionDetail.tsx
import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import { Paper, Typography, Box, CircularProgress } from '@mui/material';
import { ThemeProvider, createTheme } from '@mui/material/styles';

interface PetitionData {
  ID: number;
  title: string;
  description: string;
  pdf_url: string | null;
  signs: number;
}

const theme = createTheme();

const PetitionDetail: React.FC = () => {
  const { id } = useParams<{ id: string }>();
  const [petition, setPetition] = useState<PetitionData | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchPetition = async () => {
      try {
        const response = await fetch(`http://localhost:4242/petitions/${id}`);
        const data = await response.json();
        setPetition(data.petition);
        setLoading(false);
      } catch (error) {
        console.error('Error fetching petition:', error);
      }
    };

    fetchPetition();
  }, [id]);

  if (loading) {
    return (
      <ThemeProvider theme={theme}>
        <Box display="flex" justifyContent="center" alignItems="center" height="100vh">
          <CircularProgress />
        </Box>
      </ThemeProvider>
    );
  }

  return (
    <ThemeProvider theme={theme}>
      <Paper elevation={3} sx={{ maxWidth: 600, margin: '20px auto', padding: '20px' }}>
        <Typography variant="h4" gutterBottom>
          {petition?.title}
        </Typography>
        <Typography variant="body1" gutterBottom>
          {petition?.description}
        </Typography>
        {petition?.pdf_url && (
          <Box mt={2}>
            <a href={petition.pdf_url} target="_blank" rel="noopener noreferrer">
              View PDF
            </a>
          </Box>
        )}
        <Typography variant="body2" color="textSecondary">
          Signatures: {petition?.signs}
        </Typography>
      </Paper>
    </ThemeProvider>
  );
};

export default PetitionDetail;
